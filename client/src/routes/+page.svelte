<svelte:head>
	<link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
	  integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
	  crossorigin=""/>
</svelte:head>

<script lang="ts">
	import L from 'leaflet' // Imports Leaflet library as 'L'.

	const coords = [[51.5, -0.09], [51.49, -0.08], [51.48, -0.08]] // Unused coordinates.
	const coords2 = [[51.48, -0.09], [51.485, -0.112], [51.495, -0.115]] // Also unused coordinates.

	async function initMap(node:HTMLElement) {
		// Initializes map inside the given HTML 'node'.
		const map = L.map(node).setView([51.485, -0.095], 3);

		// Adds OpenStreetMap base tiles to the map.
		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
		}).addTo(map);

		// var testing:L.LatLngExpression[] = [
		// 	[51.509, -0.08],
		// 	[51.503, -0.06],
		// 	[51.51, -0.047]
		// ]

		// L.polygon(testing).addTo(map);
	}

	async function createPolygon() {
    var coordinates = await fetch("http://localhost:8080/extinctatlas/map").then(res => res.json())
		console.log(coordinates.length)
		for (let i = 0; i < coordinates.length; i++) {
			console.log(coordinates[i].Coordinates)
		}
  }

	createPolygon()
</script>

<div class="h-[100vh]" use:initMap></div>