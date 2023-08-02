<script>
    import Card from "./Card.svelte"
    import { Modals, closeModal, openModal } from "svelte-modals"

    export let minifyr

    async function updateMinifyr(data) {
        const minifyrJson = {
            redirect: data.redirect,
            minifyr: data.minifyr,
            random: data.random,
            id: data.id
        }
    }

</script>

<Card>
    <p>Minifyr: http://localhost:3000/r/ {minifyr.minifyr}</p>
    <p>Redirect: {minifyr.redirect}</p>
    <p>Clicked: {minifyr.clicked}</p>
    <button class="update">Update</button>
    <button class="delete">Delete</button>
</Card>

<style>
    button {
        color: white;
        font-weight: bolder;
        border: none;
        padding: .75rem;
        border-radius: 5px;
    }
    .update {
        background-color: green;
    }
    .delete {
        background-color: red;
    }
</style>
