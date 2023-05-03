<script lang="ts">
    import { link } from "svelte-navigator";
    import { token } from "../store";
    import { useLocation } from "svelte-navigator";
    import { GetProduct } from "./requests/get";
    const location = useLocation()
    let search:string

    location.subscribe(async (value)=>{
        search = value.search;
    })
</script>

<nav>
    <!-- <Link to="/">Home</Link> -->
    <h1 class="logo">
        <span>酒</span>の質
        <div class="sub-info">Sake no Shitsu</div>
    </h1>
   
    {#if $location.pathname != "/signin"}
    <div class="navigation">
        <a href="/product" class="link" class:active={search=="" && $location.pathname == "/product" } use:link>All Liquor</a>
        <a href="/product?whiskey" class="link" class:active={search=="?whiskey"} use:link>Whiskey</a>
        <a href="/product?gin" class="link" class:active={search=="?gin"} use:link>Gin</a>
        <a href="/product?vodka" class="link" class:active={search=="?vodka"} use:link>Vodka</a>
        <a href="/product?tequila" class="link" class:active={search=="?tequila"} use:link>Tequila</a>
    </div>
    

    <div class="right">
        <a href="/" use:link><svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 96 960 960" width="24"><path d="M784 936 532 684q-30 24-69 38t-83 14q-109 0-184.5-75.5T120 476q0-109 75.5-184.5T380 216q109 0 184.5 75.5T640 476q0 44-14 83t-38 69l252 252-56 56ZM380 656q75 0 127.5-52.5T560 476q0-75-52.5-127.5T380 296q-75 0-127.5 52.5T200 476q0 75 52.5 127.5T380 656Z"/></svg></a>
        <a href="/cart" use:link>
            <svg
                xmlns="http://www.w3.org/2000/svg"
                height="24"
                viewBox="0 96 960 960"
                width="24"
                ><path
                    d="M280 976q-33 0-56.5-23.5T200 896q0-33 23.5-56.5T280 816q33 0 56.5 23.5T360 896q0 33-23.5 56.5T280 976Zm400 0q-33 0-56.5-23.5T600 896q0-33 23.5-56.5T680 816q33 0 56.5 23.5T760 896q0 33-23.5 56.5T680 976ZM246 336l96 200h280l110-200H246Zm-38-80h590q23 0 35 20.5t1 41.5L692 574q-11 20-29.5 31T622 616H324l-44 80h480v80H280q-45 0-68-39.5t-2-78.5l54-98-144-304H40v-80h130l38 80Zm134 280h280-280Z"
                /></svg
            >
        </a>

        {#if $token}
            <a href="/" use:link>
                <svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 96 960 960" width="24"><path d="M480 576q-66 0-113-47t-47-113q0-66 47-113t113-47q66 0 113 47t47 113q0 66-47 113t-113 47ZM160 896V784q0-34 17.5-62.5T224 678q62-31 126-46.5T480 616q66 0 130 15.5T736 678q29 15 46.5 43.5T800 784v112H160Zm80-80h480v-32q0-11-5.5-20T700 750q-54-27-109-40.5T480 696q-56 0-111 13.5T260 750q-9 5-14.5 14t-5.5 20v32Zm240-320q33 0 56.5-23.5T560 416q0-33-23.5-56.5T480 336q-33 0-56.5 23.5T400 416q0 33 23.5 56.5T480 496Zm0-80Zm0 400Z"/></svg>
               </a
            >
        {:else}
            <a href="/signin" class="signin" use:link><h4>Sign in</h4></a>
        {/if}
    </div>
    {/if}
</nav>

<style>
    nav {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 100%;
        padding: 10px 20px;
    }
    .logo {
        text-transform: uppercase;
        font-weight: 300;
        line-height: 1;
    }
    .logo span {
        font-weight: 700;
    }
    .logo .sub-info {
        font-size: 12px;
        letter-spacing: 4px;
    }
    .navigation a {
        text-decoration: none;
        color: black;
        text-transform: uppercase;
        font-size: 16px;
        margin-left: 20px;
        font-weight: 700;
    }
    .active{
        color: #ffd700 !important;
    }
    .right {
        display: flex;
        align-items: center;
    }
   
    .right a{
        margin-left: 5px;
    }
    .right a svg{
        margin-bottom: -5px;
    }
    .signin{
        outline: none;
        text-decoration: none;
        color: black;
        border: 1px solid rgba(0, 0, 0, 0.356);
        border-radius: 10px;
        font-size: 15px;
        padding: 5px 10px;
    }
    .signin:focus{
        border:1px solid black;
    }

</style>
