<svelte:head>
	<link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
	  integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
	  crossorigin=""/>
</svelte:head>

<script lang="ts">
	import L from 'leaflet' // Imports Leaflet library as 'L'.

	const coords = [[51.5, -0.09], [51.49, -0.08], [51.48, -0.08]] // Unused coordinates.
	const coords2 = [[51.48, -0.09], [51.485, -0.112], [51.495, -0.115]] // Also unused coordinates.

	var cooResult:Array<L.LatLngExpression[]> = [];

	async function getPolygonInfo():Promise<Array<L.LatLngExpression[]>> {
    var coordinates = await fetch("http://localhost:8080/extinctatlas/map").then(res => res.json())

		for (let i = 0; i < coordinates.length; i++) {
			let gatherer:L.LatLngExpression[] = []
			for (let index = 0; index < coordinates[i].Coordinates.length; index++) {
				gatherer.push([coordinates[i].Coordinates[index].Lat, coordinates[i].Coordinates[index].Lng])
			}
			cooResult.push(gatherer)
		}
		return cooResult
  }
	


	async function initMap(node:HTMLElement) {
		// Initializes map inside the given HTML 'node'.
		const map = L.map(node).setView([51.485, -0.095], 3);

		// Adds OpenStreetMap base tiles to the map.
		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
		}).addTo(map);

		var testing:L.LatLngExpression[] = [
			[51.509, -0.08],
			[51.503, -0.06],
			[51.51, -0.047]
		]
		await getPolygonInfo().then(res => {
			cooResult.map(el => {
				L.polygon(el).addTo(map);
			})
		})
	}
</script>

<div class="h-[100vh]" use:initMap></div>