<div class="loader">
    <div>
        <div class="wineglass left">
            <div class="top" />
        </div>
        <div class="wineglass right">
            <div class="top" />
        </div>
    </div>
</div>

<style>
    .loader {
        position: fixed;
        width: 100vw;
        height: 100vh;
        z-index: 1000;
        background-color: white;
        display: grid;
        place-items: center;
    }
    .wineglass {
        height: 30px;
        position: relative;
    }
    .wineglass:after {
        content: "";
        position: absolute;
        top: 40px;
        left: 3px;
        width: 20px;
        height: 3px;
        background: black;
        box-shadow: 0 0 1px black;
    }

    .top {
        background: black;
        width: 25px;
        height: 30px;
        border-radius: 0 0 36px 36px;
        box-shadow: 0 0 1px black;
    }
    .top:before {
        content: "";
        position: absolute;
        left: 2.4px;
        bottom: 3px;
        width: 20px;
        height: 22px;
        border-radius: 50%;
        background: linear-gradient(transparent 50%, white 50%);
    }
    .top:after {
        content: "";
        position: absolute;
        top: 30px;
        left: 11px;
        width: 3px;
        height: 13px;
        background: black;
        box-shadow: 0 0 1px black;
    }

    .left {
        display: inline-block;
        margin-right: 10px;
        -webkit-animation: rotate1 2s cubic-bezier(0.39, 1.52, 0.46, 0.92)
            infinite;
        animation: rotate1 2s cubic-bezier(0.39, 1.52, 0.46, 0.92) infinite;
    }
    .left .top:before {
        -webkit-animation: rotate2 2s linear infinite;
        animation: rotate2 2s linear infinite;
    }

    .right {
        display: inline-block;
        -webkit-animation: rotate2 2s cubic-bezier(0.39, 1.52, 0.46, 0.92)
            infinite;
        animation: rotate2 2s cubic-bezier(0.39, 1.52, 0.46, 0.92) infinite;
    }
    .right .top:before {
        -webkit-animation: rotate1 2s linear infinite;
        animation: rotate1 2s linear infinite;
    }

    @-webkit-keyframes rotate1 {
        0% {
            transform: rotate(0deg);
        }
        50% {
            transform: rotate(22deg);
        }
    }

    @keyframes rotate1 {
        0% {
            transform: rotate(0deg);
        }
        50% {
            transform: rotate(22deg);
        }
    }
    @-webkit-keyframes rotate2 {
        0% {
            transform: rotate(0deg);
        }
        50% {
            transform: rotate(-22deg);
        }
    }
    @keyframes rotate2 {
        0% {
            transform: rotate(0deg);
        }
        50% {
            transform: rotate(-22deg);
        }
    }
</style>
