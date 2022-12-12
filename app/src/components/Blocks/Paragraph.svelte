<script lang="ts">
    import {createEventDispatcher} from 'svelte';
    import {getCaretOffset,isCaretAtStart} from "../../lib/contentEditablePos.js"

    const dispatch = createEventDispatcher();


    export let focus = false;
    let focused = false;

    let input = null;
    export let content = ""

    $: {
        if (focus && input != null) {
            focusInput();
        }

    }

    function keyBinding(evt) {

        if (evt.key === "Enter" && !evt.shiftKey) {
            console.log("pass here");
            evt.preventDefault();
        }

        setTimeout(() => {
            const cursorPosition = getCaretOffset(input);
            console.log(cursorPosition)

            dispatch("change", content)
            switch (evt.key) {
                case "ArrowUp":
                    dispatch("prev", cursorPosition)
                    break;
                case "ArrowDown":
                    dispatch("next", cursorPosition)
                    break;
                case "Backspace":
                    dispatch("remove.prev", isCaretAtStart(input));
                    break;
                case "Enter":
                    if (!evt.shiftKey) {
                        dispatch("add.next", cursorPosition);
                        console.log(content);
                    }
                    break;
            }


        }, 10);


    }

    function onBlur() {
        focused = false;
        dispatch("blur")
    }

    function onFocus() {


        focused = true

        dispatch("focus")
    }

    function focusInput() {
        input.focus();
    }

</script>


<div class="" on:focus={focusInput}>

    {#if content.length === 0}
        {#if focused}
            <div class="text-gray-200 p-2 absolute z-0 select-none">Type '/' for commands</div>
        {/if}
    {/if}

    <div class="w-full whitespace-pre-wrap	 z-10 p-2 ring-0 ring-offset-0 ring-transparent outline-0	"
         contenteditable="true"
         bind:this={input}
         on:blur={onBlur}
         on:focus={onFocus}
         on:keydown={keyBinding}
         bind:innerHTML={content}/>
</div>
