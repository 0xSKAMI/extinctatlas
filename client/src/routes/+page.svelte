<svelte:head>
  <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
    integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
    crossorigin=""/>
</svelte:head>


<script lang="ts">
  import L from 'leaflet' // Imports Leaflet library as 'L'.


  var info:object = {}
  var AiInput:string = ""
	var id:string = ""
	var aiResponding = false;
	var response = "";
	var loading = false 


  var cooResult:Array<L.LatLngExpression[]> = [];


  async function getPolygonInfo():Promise<[Array<L.LatLngExpression[]>, Array<object>]> {
    var coordinates = await fetch("http://localhost:8080/extinctatlas/map").then(res => res.json())


    for (let i = 0; i < coordinates.length; i++) {
      let gatherer:L.LatLngExpression[] = []
      for (let index = 0; index < coordinates[i].Coordinates.length; index++) {
        gatherer.push([coordinates[i].Coordinates[index].Lat, coordinates[i].Coordinates[index].Lng])
      }
      cooResult.push(gatherer)
    }
    return [cooResult, coordinates]
  }


  async function getInfo(id: string) {
    info = await fetch('http://localhost:8080/extinctatlas/info/' + id).then(res => res.json())
  }
 
  async function getAiResponse(prompt: string, id: string) {
		loading = true;
		await fetch('http://localhost:8080/extinctatlas/ai?id='+id+'&prompt='+prompt).then(res => res.text()).then(res => {response = res; loading = false})
  }


  async function initMap(node:HTMLElement) {
    // Initializes map inside the given HTML 'node'.
    const map = L.map(node).setView([51.485, -0.095], 3);


    // Adds OpenStreetMap base tiles to the map.
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    await getPolygonInfo().then(([res, arr]) => {
      cooResult.map((el, index) => {
        L.polygon(el).addTo(map).addEventListener("click", async() => {await getInfo(arr[index].ID); id = arr[index].ID});
      })
    })
  }
</script>


<div class="flex">
  {#if Object.keys(info).length > 0}
  <div class="h-screen w-64 bg-slate-800 text-gray-200 p-6 shadow-lg flex flex-col space-y-4">
		<button class="self-end px-3 py-1 bg-slate-700 hover:bg-slate-600 rounded text-sm" onclick={() => {info = {}; id = "", response = "", aiResponding = false}}>close</button>
		{#if aiResponding == false}
			<img class="object-cover rounded-md self-center border-2 border-slate-600" src="{info.ImageURL}" alt="{info.Name}">
			<div class="text-center">
				<p class="text-xl font-semibold">{info.Name}</p>
				<p class="text-sm text-gray-400">{info.Type}</p>
			</div>
			<div>
				<p><span class="font-semibold text-gray-300">Reason of extinction:</span> {info.Reason}</p>
				<p><span class="font-semibold text-gray-300">Last seen:</span> {info.LastSeen}</p>
				<p><span class="font-semibold text-gray-300">Height:</span> {info.HeightCM}cm</p>
				<p><span class="font-semibold text-gray-300">Weight:</span> {info.WeightKG}kg</p>
				<p><span class="font-semibold text-gray-300">Diet: </span>
				{#each info.Diet as meal}
					{" " + meal}
				{/each}</p>
			</div>
		{/if}
		{#if aiResponding == true}
			{#if loading == true}
				<div class="flex-col gap-4 w-full flex items-center justify-center">
					<div class="w-20 h-20 border-4 border-transparent text-blue-400 text-4xl animate-spin flex items-center justify-center border-t-blue-400 rounded-full">
						<div class="w-16 h-16 border-4 border-transparent text-red-400 text-2xl animate-spin flex items-center justify-center border-t-red-400 rounded-full"></div>
					</div>
				</div>
			{/if}
			<p>{response}</p>
		{/if}
    <div class="mt-auto">
      <input
        bind:value={AiInput}
				onkeydown="{async (e) => {if (e.key == "Enter" && AiInput != "") {response = ""; aiResponding = true; await getAiResponse(AiInput, id); AiInput = ""}}}"
        class="bg-[#222630] px-4 py-3 outline-none w-full text-white rounded-lg border-2 transition-colors duration-100 border-solid focus:border-[#596A95] border-[#2B3040]"
        name="text"
        placeholder="Ask ai"
        type="text"
      />
    </div>
  </div>
  {/if}
  <div class="h-[100vh] w-screen" use:initMap></div>
</div>
