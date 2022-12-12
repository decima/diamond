<script>
    import Paragraph from "../Blocks/Paragraph.svelte";

    let currentBlock = 0;
    let cursorPosition = 0;
    let contents = [
        {
            text: "a"
        },
        {
            text: "b"
        },
        {
            text: "c"
        }
    ];

    function goNext(currentPosition) {
        currentBlock++;
    }

    function goPrev(currentPosition) {
        currentBlock--;
    }

    function addNext(currentPosition) {
        let nc = ""
        if (currentPosition !== null) {
            console.log(currentPosition)
            nc = contents[currentBlock].text.substring(currentPosition);
            contents[currentBlock].text = contents[currentBlock].text.substring(0, currentPosition);
        }
        contents.splice(currentBlock + 1, 0, {text: nc})
        goNext(0);
    }

    function removePrev(caretAtFirst) {
        if (!caretAtFirst) {
            return;
        }

        let nc = contents[currentBlock].text
        contents[currentBlock - 1].text += nc

        goPrev(0);
        contents.splice(currentBlock + 1, 1)


    }


</script>
{currentBlock}
<div className="flex">
    <div className="flex w-1/2 flex-col bg-slate-100">
        {#each contents as content,i }
            <Paragraph focus={currentBlock===i}
                       focusAt={cursorPosition}
                       content={content.text}
                       on:focus={()=>currentBlock=i}
                       on:next={(e)=>goNext(e.detail)}
                       on:prev={(e)=>goPrev(e.detail)}
                       on:add.next={(e)=>addNext(e.detail)}
                       on:remove.prev={removePrev}
                       on:change={(e)=>{contents[i].text=e.detail}}
                       on:add={()=>{console.log("perfect");}}
            />
        {/each}
    </div>
    <pre class="w-1/2">{JSON.stringify(contents, null, 4)}</pre>
</div>