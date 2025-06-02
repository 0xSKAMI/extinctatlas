<svelte:head>
	<link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.4/dist/leaflet.css"
	  integrity="sha256-p4NxAoJBhIIN+hmNHrzRCf9tD/miZyoHS5obTRR9BMY="
	  crossorigin=""/>
</svelte:head>

<script lang="ts">
	import L from 'leaflet' // Imports Leaflet library as 'L'.

	const coords = [[51.5, -0.09], [51.49, -0.08], [51.48, -0.08]] // Unused coordinates.
	const coords2 = [[51.48, -0.09], [51.485, -0.112], [51.495, -0.115]] // Also unused coordinates.

	function initMap(node:HTMLElement) {
		// Initializes map inside the given HTML 'node'.
		const map = L.map(node).setView([51.485, -0.095], 13);

		// Adds OpenStreetMap base tiles to the map.
		L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
			attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
		}).addTo(map);

		var polygon = L.polygon([
			[51.509, -0.08],
			[51.503, -0.06],
			[51.51, -0.047]
		]).addTo(map); // Draws a polygon on the map.

		polygon.bindPopup("Hello") // Adds "Hello" popup on polygon click.
	}
</script>

<div class="h-[100vh]" use:initMap></div>