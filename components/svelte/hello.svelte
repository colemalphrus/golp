<script>
    var greeting = "This page is Svelte Powered"
    $: count = 0
</script>

<div>
    <h1>{greeting}</h1>
    <h2>{count}</h2>
    <button on:click={() => count += 1}>ADD ONE</button>
    <button on:click={() => count -= 1}>Sub ONE</button>
    <button on:click={() => count = 0}>CLEAR</button>
</div>

<style>
    div {
        background-color: cadetblue;
        text-align: center;
    }
    button {
        width: 8rem;
        height: 4rem;
        border-radius: 10px;

    }
</style>