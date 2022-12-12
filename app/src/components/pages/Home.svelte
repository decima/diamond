<script>

    function uid() {
        return Date.now().toString(36) + Math.random().toString(36).substr(2);
    }

    let contents = [
        {
            tag: "p",
            html: ""
        }
    ];


</script>
{currentBlock}
<div className="flex">
    <div className="flex w-1/2 flex-col bg-slate-100">
        {#each contents as content,i }

        {/each}
    </div>
    <pre class="w-1/2">{JSON.stringify(contents, null, 4)}</pre>
</div>