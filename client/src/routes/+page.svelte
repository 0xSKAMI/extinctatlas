<script lang="ts">
  import { text } from '@sveltejs/kit';
import { MapLibre } from 'svelte-maplibre-gl';
</script>

<MapLibre 
  class="h-[100vh]"
  style="https://basemaps.cartocdn.com/gl/voyager-gl-style/style.json"
  >
  
</MapLibre>

<style>
  @reference "tailwindcss";
</style>