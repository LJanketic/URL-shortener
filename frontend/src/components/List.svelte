<script>
    import Item from "./Item.svelte"
    import { onMount } from "svelte"

    let minifyrs = []
    onMount(async() => {
        const response = await fetch("http://localhost:3000/minifyr")
        minifyrs = await response.json()
    })
</script>

{#each minifyrs as minifyr}
    <Item minifyr={minifyr} />
{/each}
