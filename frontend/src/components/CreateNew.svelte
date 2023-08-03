<script>
    import { Modals, closeModal, openModal } from "svelte-modals"
    import Modal from "./Modal.svelte";

    async function createMinifyr(data) {
        const minifyrJson = {
            redirect: data.redirect,
            minifyr: data.minifyr,
            random: data.random
        }
        await fetch("http://localhost:3000/minifyr", {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(minifyrJson)
        }).then(response => {
            console.log(response)
        })
    }

    function handleOpenModal() {
        openModal(Modal, {
            title: "Create new minifyr",
            send: createMinifyr,
            redirect: "",
            minifyr: "",
            random: false
        })
    }
</script>

<button on:click={handleOpenModal}>Add New</button>

<style>
    button {
        background-color: green;
        color: white;
        font-weight: bold;
        border: none;
        padding: .75rem;
        border-radius: 4px;
    }
</style>
