<script lang="ts">
    import { Alerts } from "../store";
    import { fade } from 'svelte/transition';
    function False() {
        Alerts.update((value) => {
            return { ...value, show: false };
        });
    }

    function HandleOk() {
        False();
        $Alerts.handleOk();
    }

    function HandleCancel() {
        False();
        $Alerts.handleCancel();
    }
</script>

{#if $Alerts.show}
    <div class="alertscreen noselect" in:fade={{duration:100}} out:fade={{duration:50}}>
        <div class="notifications-container">
            <div class="alert">
                <div class="flex">
                    <div class="alert-prompt-wrap">
                        <p class="alert-prompt-heading">{$Alerts.title}</p>
                        <div class="alert-prompt-prompt">
                            <p>{$Alerts.message}</p>
                        </div>
                        <div class="alert-button-container">
                            <button
                                type="button"
                                class="alert-button-main"
                                on:click={HandleOk}>Ok</button
                            >
                            <button
                                type="button"
                                class="alert-button-secondary"
                                on:click={HandleCancel}>Cancel</button
                            >
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
{/if}

<style>
    .alertscreen {
        position: fixed;
        width: 100%;
        height: 100%;
        z-index: 1000;
        background-color: rgba(0, 0, 0, 0.753);
        display: grid;
        place-items: center;
    }
    

    .flex {
        display: flex;
    }

    .alert {
        padding: 1rem;
        border-radius: 0.375rem;
        background-color: rgb(255, 255, 255);
    }

    .alert-prompt-wrap {
        margin-left: 0.75rem;
    }

    .alert-prompt-heading {
        font-weight: bold;
        color: var(--primary-color);
    }

    .alert-prompt-prompt {
        margin-top: 0.5rem;
        font-weight: 500;
        color: var(--primary-color);
    }

    .alert-button-container {
        display: flex;
        margin-top: 0.875rem;
        margin-bottom: -0.375rem;
        margin-left: -0.5rem;
        margin-right: -0.5rem;
    }

    .alert-button-main {
        padding-top: 0.375rem;
        padding-bottom: 0.375rem;
        padding-left: 0.5rem;
        padding-right: 0.5rem;
        background-color: #f8f8f8;
        font-size: 0.875rem;
        line-height: 1.25rem;
        font-weight: bold;
        border-radius: 0.375rem;
        border: none;
    }

    .alert-button-main:hover {
        background-color: #e3e3e3;
    }

    .alert-button-secondary {
        padding-top: 0.375rem;
        padding-bottom: 0.375rem;
        padding-left: 0.5rem;
        padding-right: 0.5rem;
        margin-left: 0.75rem;
        background-color: transparent;
        color: #020202;
        font-size: 0.875rem;
        line-height: 1.25rem;
        border-radius: 0.375rem;
        border: none;
    }
</style>
