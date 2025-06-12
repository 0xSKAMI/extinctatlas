<svelte:head>
  <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
    integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
    crossorigin=""/>
</svelte:head>


<script lang="ts">
  import L from 'leaflet'
	import { spring } from "svelte/motion"


  var info:object = {}
  var AiInput:string = ""
	var id:string = ""
	var aiResponding = false;
	var response = "";
	var loading = false

  var cooResult:Array<L.LatLngExpression[]> = [];

  async function getPolygonInfo():Promise<[Array<L.LatLngExpression[]>, Array<object>]> {
    var coordinates = await fetch("https://extinctatlas.duckdns.org/extinctatlas/map").then(res => res.json())


    for (let i = 0; i < coordinates.length; i++) {
      let gatherer:L.LatLngExpression[] = []
      for (let index = 0; index < coordinates[i].Coordinates.length; index++) {
        gatherer.push([coordinates[i].Coordinates[index].Lat, coordinates[i].Coordinates[index].Lng])
      }
      cooResult.push(gatherer)
    }
    return [cooResult, coordinates]
  }


  async function getInfo(id_param: string) {
    info = await fetch('https://extinctatlas.duckdns.org/extinctatlas/info/' + id_param).then(res => res.json())
    id = id_param;
  }

  async function getAiResponse(prompt: string, animal_id: string) {
		loading = true;
		await fetch('https://extinctatlas.duckdns.org/extinctatlas/ai?id='+animal_id+'&prompt='+prompt).then(res => res.text()).then(res_text => {response = res_text; loading = false})
  }


  async function initMap(node:HTMLElement) {
    const map = L.map(node).setView([51.485, -0.095], 3);


    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    await getPolygonInfo().then(([res, arr]) => {
      cooResult.map((el, index) => {
				var polygon = L.polygon(el).addTo(map);
				polygon.addEventListener("click", async() => {await getInfo(arr[index].ID); response = ""; aiResponding = false});
				polygon.bindPopup('<div class="w-48 font-sans antialiased">' + '<img class="w-full rounded-lg object-cover shadow-sm mb-2" src="' + arr[index].ImageURL + '"/>' + '<p class="text-center text-sm font-semibold text-slate-700 leading-tight break-words px-2 py-1">' + arr[index].Name + '</p>' + '</div>');
				polygon.on("mouseover", (e) => {
					polygon.openPopup();
				})
				polygon.on("mouseout", (e) => {
					polygon.closePopup();
				})
      })
    })
  }
</script>


<div class="flex relative">
  {#if Object.keys(info).length > 0}
  <div class="bg-slate-800 text-gray-200 flex flex-col z-[1000]
              fixed bottom-0 left-0 right-0 max-h-[65vh] p-4 space-y-3 rounded-t-xl shadow-xl overflow-y-auto
              md:static md:h-screen md:w-64 md:max-h-none md:p-6 md:space-y-4 md:rounded-none md:shadow-lg">
		<button class="self-end px-3 py-1 bg-slate-700 hover:bg-slate-600 rounded text-sm" onclick={() => {info = {}; id = "", response = "", aiResponding = false}}>close</button>
		{#if aiResponding == false}
			<img class="object-cover rounded-md self-center border-2 border-slate-600 max-w-[80%] md:max-w-full" src="{info.ImageURL}" alt="{info.Name}">
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
				<div class="flex-col gap-4 w-full flex items-center justify-center py-5">
					<div class="w-20 h-20 border-4 border-transparent text-blue-400 text-4xl animate-spin flex items-center justify-center border-t-blue-400 rounded-full">
						<div class="w-16 h-16 border-4 border-transparent text-red-400 text-2xl animate-spin flex items-center justify-center border-t-red-400 rounded-full"></div>
					</div>
				</div>
			{/if}
			{#if response && response.trim() !== ""}
				<div class="overflow-y-auto flex-1 min-h-0">
					<p class="whitespace-pre-wrap text-sm md:text-base">{response}</p>
				</div>
			{/if}
		{/if}
    <div class="mt-auto pt-2">
      <input
        bind:value={AiInput}
				onkeydown="{async (e) => {if (e.key == "Enter" && AiInput.trim() != "") {response = ""; aiResponding = true; await getAiResponse(AiInput, id); AiInput = ""}}}"
        class="bg-[#222630] px-4 py-3 outline-none w-full text-white rounded-lg border-2 transition-colors duration-100 border-solid focus:border-[#596A95] border-[#2B3040]"
        name="text"
        placeholder="Ask ai"
        type="text"
      />
    </div>
  </div>
  {/if}
  <div class="h-screen w-full md:flex-1" use:initMap></div>
</div>
