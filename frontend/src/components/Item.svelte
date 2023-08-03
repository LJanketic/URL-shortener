<script>
    import Card from "./Card.svelte"
    import Modal from "./Modal.svelte"
    import { Modals, closeModal, openModal } from "svelte-modals"

    export let minifyr
    let displayCard = true

    async function updateMinifyr(data) {
        const minifyrJson = {
            redirect: data.redirect,
            minifyr: data.minifyr,
            random: data.random,
            id: data.id
        }
        await fetch("gttp://localhost:3000/redirect", {
            method: "PATCH",
            headers: {"Content-type": "application/json"},
            body: JSON.stringify(minifyrJson)
        }).then(res => {
            displayCard = false
            console.log(res)
        })
    }

    function handleOpenModal(minifyr){
        openModal(Modal, {
            title: "Update Minifyr link",
            send: updateMinifyr,
            minifyr: minifyr.minifyr,
            redirect: minifyr.redirect,
            random: minifyr.random
        })
    }

    async function deleteMinifyr(){
        if(confirm("Are you certain you wish to delete the following Minifyr link (" + minifyr.minifyr + ")?")){
            await fetch("http://localhost:3000/minifyr/" + minifyr.id, {
                method: "DELETE"
            }).then(response => {
            console.log(response)
        })
        }
    }

</script>
{#if displayCard}
<Card>
    <p>Minifyr: http://localhost:3000/r/{minifyr.minifyr}</p>
    <p>Redirect: {minifyr.redirect}</p>
    <p>Clicked: {minifyr.clicked}</p>
    <button class="update" on:click={handleOpenModal(minifyr)}>Update</button>
    <button class="delete" on:click={deleteMinifyr}>Delete</button>
</Card>
{/if}
<Modals>
    <div
        slot="backdrop"
        class="backdrop"
        on:click={closeModal}
    />
</Modals>

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
    .backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        right: 0;
        left: 0;
        background: rgb(255, 255, 255)
    }
</style>
