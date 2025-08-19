<!-- for leaflet js -->
<svelte:head>
  <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
    integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
    crossorigin=""/>
</svelte:head>


<!-- starting of typescript part -->
<script lang="ts">
	//importing leaflet
  import L from 'leaflet'
	//for animations
  import { fly } from "svelte/transition";
	//for animations
	import { spring, tweened } from "svelte/motion"
	//for .env
	import { PUBLIC_ADDRESS } from "$env/static/public";

	const mouseCoords = spring({ x: 0, y: 0 })

	const onMouseMove = (event) => {
		$mouseCoords = { x: event.x, y: event.y }
	}

  var info:object = {}
  var AiInput:string = ""
	var id:string = ""
	var aiResponding = false;
	var response = "";
	var loading = false
  var questions = []
	var popupImgUrl = ""
	var popupName = ""
	var popupOpacity = 0;

  var cooResult:Array<L.LatLngExpression[]> = [];

	//function fo get polygon coordinates
  async function getPolygonInfo():Promise<[Array<L.LatLngExpression[]>, Array<object>]> {
		//calling api enpoint and turniong result in json
    var coordinates = await fetch(PUBLIC_ADDRESS + "/map").then(res => res.json())

		//doing loop to to create array of points
    for (let i = 0; i < coordinates.length; i++) {
      cooResult.push(coordinates[i].Polygon.coordinates)
    }
    return [cooResult, coordinates]
  }


	//function to get info about single animal
  async function getInfo(id_param: string) {
    info = await fetch(PUBLIC_ADDRESS+'/info/' + id_param).then(res => res.json())
    id = id_param;
  }

	//function to ai answer to question 
  async function getAiResponse(prompt: string, animal_id: string) {
		loading = true;
		await fetch(PUBLIC_ADDRESS+'/ai?id='+animal_id+'&prompt='+prompt).then(res => res.text()).then(res_text => {response = res_text; loading = false})
  }
  
	//function to AI generated example questions
  async function getAiQuesions(animal_id: string) {
		loading = true;
		await fetch(PUBLIC_ADDRESS+'/ai/generate?id='+animal_id).then(res => res.text()).then(res_text => questions = res_text.split("/"))
  }

	//funciton to generate map and map content
  async function initMap(node:HTMLElement) {
		//declare map and set starting point
    const map = L.map(node).setView([51.485, -0.095], 3);

		//configure map basically 
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

		//getting polygon coordinates and appying them to map
    await getPolygonInfo().then(([res, arr]) => {
				cooResult.map((el, index) => {
					var polygon = [{
						"type": "Feature",
						"properties": {"name": "animal", "image": arr[index].ImageURL},
						"geometry": {
								"type": "Polygon",
								"coordinates": 
									el
						}
				}];
				L.geoJSON(polygon, {
					onEachFeature: function(feature, layer) {
						layer.addEventListener("click", async() => {await getInfo(arr[index].ID); await getAiQuesions(arr[index].ID); response = ""; aiResponding = false});
						layer.on("mouseover", (e) => {
							popupImgUrl = arr[index].ImageURL;
							popupName = arr[index].Name;
							popupOpacity = 100;
						})
						layer.on("mouseout", (e) => {
							popupOpacity = 0;
						})
					}
				}).addTo(map);
      })
    })
  }
</script>

<svelte:window 
	on:mousemove={onMouseMove} 
/>

<div style:--x={`${$mouseCoords.x}px`} style:--y={`${$mouseCoords.y - 100}px`} style:opacity={popupOpacity / 100} class="absolute top-0 left-0 w-[10vh] z-[999] transform translate-x-[calc(-50%_+_var(--x,_0px))] translate-y-[calc(-50%_+_var(--y,_0px))] scale-[var(--scale,1)] pointer-events-none bg-white p-2 rounded-lg shadow-xl ring-1 ring-black ring-opacity-5 flex flex-col items-center transition-opacity duration-300 ease-in-out">
	<img class="w-full rounded-lg object-cover shadow-sm mb-2" alt="animal" src={popupImgUrl}/>
	<p class="text-center text-xs font-semibold text-slate-700 leading-tight break-words px-1 py-0.5"> {popupName} </p>
</div>

<div class="flex relative">
  {#if Object.keys(info).length > 0}
    <div
      transition:fly={{y: '100%', duration: 300, opacity: 1 }}
      class="bg-slate-900/80 backdrop-blur-md text-slate-300 flex flex-col z-[1000]
            fixed bottom-0 left-0 right-0 max-h-[70vh] rounded-t-2xl shadow-2xl
            md:static md:bg-slate-900 md:backdrop-blur-none md:h-screen md:w-80 md:max-h-none md:rounded-none md:shadow-lg"
    >
      <!-- Main Content Scroll Area -->
      <div class="flex-1 overflow-y-auto p-5 md:p-6">
        <button
          class="absolute top-3 right-3 z-10 flex h-8 w-8 items-center justify-center rounded-full bg-slate-800/60 text-slate-400 transition-colors hover:bg-slate-700 hover:text-slate-100
                 md:self-end md:static md:mb-4"
          on:click={() => {info = {}; id = "", response = "", aiResponding = false}}
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>

        {#if aiResponding == false}
          <div class="space-y-5" transition:fly={{ y: 20, duration: 200 }}>
            <img class="w-full rounded-lg object-cover shadow-lg" src="{info.ImageURL}" alt="{info.Name}" />
            <div class="text-center">
              <h1 class="text-2xl font-bold tracking-tight text-white">{info.Name}</h1>
              <p class="font-medium text-indigo-400">{info.Type}</p>
            </div>
            <div class="text-sm">
              <h3 class="mb-3 border-b border-slate-700 pb-2 font-semibold text-white">Details</h3>
              <dl class="grid grid-cols-[max-content,1fr] gap-x-4 gap-y-2">
                <dt class="font-semibold text-slate-400">Extinction:</dt>
                <dd>{info.Reason}</dd>
                <dt class="font-semibold text-slate-400">Last Seen:</dt>
                <dd>{info.LastSeen}</dd>
                <dt class="font-semibold text-slate-400">Height:</dt>
                <dd>{info.HeightCM} cm</dd>
                <dt class="font-semibold text-slate-400">Weight:</dt>
                <dd>{info.WeightKG} kg</dd>
                <dt class="font-semibold text-slate-400">Diet:</dt>
                <dd>
                  {#each info.Diet as meal, i}
                    {meal}{#if i < info.Diet.length - 1}, {/if}
                  {/each}
                </dd>
              </dl>
            </div>
          </div>
        {/if}

        {#if aiResponding == true}
          <div class="flex h-full flex-col justify-center">
            {#if loading == true}
              <div class="flex w-full items-center justify-center py-10" transition:fly={{ y: 20, duration: 200 }}>
                <div class="h-12 w-12 animate-spin rounded-full border-4 border-slate-700 border-t-indigo-500"></div>
              </div>
            {/if}
            {#if response && response.trim() !== ""}
              <div class="min-h-0 flex-1 overflow-y-auto" transition:fly={{ y: 20, duration: 200 }}>
                <h3 class="mb-3 font-semibold text-white">AI Response</h3>
                <div class="whitespace-pre-wrap rounded-lg bg-slate-800/50 p-4 text-sm leading-relaxed ring-1 ring-slate-700 md:text-base">
                  {response}
                </div>
              </div>
            {/if}
          </div>
        {/if}
      </div>

      <!-- Footer with Input -->
      <div class="mt-auto border-t border-slate-800 bg-slate-900/50 p-3 md:bg-slate-900">
        <div class="mb-3 flex flex-wrap justify-center gap-2 px-2">
          {#each questions as question}
            <button on:click={AiInput = question} class="rounded-full bg-slate-800 px-3 py-1.5 text-xs font-medium text-slate-300 transition-colors hover:bg-slate-700">
              {question.slice(0, 23)}{#if question.length > 23}...{/if}
            </button>
          {/each}
        </div>
        <div class="relative flex items-center">
          <input
            bind:value={AiInput}
            on:keydown="{async (e) => {if (e.key === 'Enter' && AiInput.trim() != '') {response = ''; aiResponding = true; await getAiResponse(AiInput, id); AiInput = ''}}}"
            class="w-full rounded-full border-2 border-transparent bg-slate-800 py-2.5 pl-4 pr-12 text-white outline-none transition-colors duration-200 placeholder:text-slate-500 focus:border-indigo-600 focus:bg-slate-900"
            name="text"
            placeholder="Ask about the {info.Name}..."
            type="text"
          />
          <button
            class="absolute right-2 flex h-8 w-8 items-center justify-center rounded-full bg-indigo-600 text-white transition-colors hover:bg-indigo-500 disabled:bg-slate-700 disabled:text-slate-500"
            disabled={AiInput.trim() === '' || aiResponding}
            on:click="{async () => {if (AiInput.trim() != '') {response = ''; aiResponding = true; await getAiResponse(AiInput, id); AiInput = ''}}}"
          >
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="h-5 w-5">
              <path d="M3.478 2.405a.75.75 0 00-.926.94l2.432 7.905H13.5a.75.75 0 010 1.5H4.984l-2.432 7.905a.75.75 0 00.926.94 60.519 60.519 0 0018.445-8.986.75.75 0 000-1.218A60.517 60.517 0 003.478 2.405z" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  {/if}

  <div class="h-screen w-full md:flex-1" use:initMap></div>
</div>
