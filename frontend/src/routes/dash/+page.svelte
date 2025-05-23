<script lang="ts">
    import Dashboard from '$lib/components/Dash.svelte';
	import { onMount } from 'svelte';
  
    // You could fetch user data in a real app
    export let data: {
        name: string;
    };
    console.log('name 2', name);

    onMount(() => {
        console.log('name in broswer:', name);
    })
  </script>
  
  <Dashboard name={data.name} />
  