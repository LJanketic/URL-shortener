<script>
    import Card from "./Card.svelte"
</script>

<Card>
    <p>Minifyr: http://localhost:3000/r/ {"Test"}</p>
    <p>Redirect: {"Redirect test"}</p>
    <p>Clicked: </p>
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
