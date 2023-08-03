<script>
    import { Modals, closeModal, openModal } from "svelte-modals"
    import Modal from "./Modal.svelte";

    async function createMinifyr(data) {
        minifyrJson = {
            redirect: data.redirect,
            minifyr: data.minifyr,
            random: data.random
        }
        await fetch("http://localhost:3000/redirect", {
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
