<script>
    import { fade, fly } from 'svelte/transition';
    var greeting = "This about us page is made with svelte"
    var showReasons = false
    $: reasons = [
        {id: 1, r: "Go and svelte are compiled so this is fast"},
        {id: 2, r: "Every GOLP app can contain a separate svelte app"},
        {id: 3, r: "You can use mustache when you dont need better SEO"},
        {id: 4, r: "Your making landing pages not a blog so this wins"},
    ]
</script>

<div>
    <button on:click={()=> showReasons = !showReasons} >DO NOT CLICK</button>
    {#if !showReasons}
    <h1 in:fly={{ y: -1000, duration: 1000 }} out:fly={{x: -1000, duration: 1000}}>
        {greeting}
    </h1>
    <h2 in:fly={{ y: 1000, duration: 1000 }} out:fly={{x: 1000, duration: 1000}}>
        For no good reason (but its cool right?)
    </h2>
    {/if}
    <ul>
        {#if showReasons}
        {#each reasons as reason (reason.id) }
            <li 
            in:fade={{ y: -1000, duration: 2000, delay: 1000 }}
            out:fade={{x: -1000, duration: 2000, delay: 1000}}
            >
            <strong>{reason.id}</strong> {reason.r}
            </li>
        {/each}
        {/if}
    </ul>
    <div>
        <a href="/">back to home</a>
    </div>
</div>

<style>
    div {
        background-color: cadetblue;
        text-align: center;
    }
    ul {
        width: 50%;
        margin: 0 auto
    }
    li {
        text-align: left;
        margin: 1em;
        font-size: x-large;
    }
    button {
        position: absolute;
        bottom: 0;
        right: 0;
        padding: 2em;
        color: white;
        background-color: red;
        border-radius: 100%;
        height: 100px;
        width: 100px;
    }

</style>