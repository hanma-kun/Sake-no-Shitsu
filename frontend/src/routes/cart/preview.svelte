<script lang="ts">
    import { navigate, useLocation } from "svelte-navigator";
    import {  Getdetials } from "../../utils/requests/get";
    import Loader from "../../utils/loader.svelte";
    import { token } from "../../store";
    import { AddtoCart } from "../../utils/requests/post";
    let location = useLocation();
    let loading = false;
    let search = "";
    let key:string
    let mainProduct = {
        image: "",
        name: "",
        amount: 0,
        rating: 0,
    };

    

    location.subscribe(async (value) => {
        search = value.state.data;
    });

    async function Init() {
        if ($location.pathname == "/preview") {
            if (search && search != "") {
                loading = true;
                let data = await Getdetials(search);
                mainProduct.image = data.image;
                mainProduct.amount = data.price;
                mainProduct.name = data.product_name;
                mainProduct.rating = data.rating;
                loading = false;
            }
        }
    }
    Init();
    token.subscribe((value)=>{
        key = value;
    })
    async function Addtocart(){
        if(key && key !=""){
            await AddtoCart(search)
        }else{
            navigate("/signin",{replace:true})
        }
    }

</script>

{#if loading}
    <Loader />
{/if}

<div class="maincontainer">
    <div class="detialsContainer">
        <div class="card1">
            <div class="imagecon">
                <img src={mainProduct.image} alt="" />
            </div>
            <div class="detialscon">
                <div>
                    <h1>{mainProduct.name}</h1>
                </div>
                <div>
                    <div style="display: flex;">
                        <h3>Price</h3>
                        <span style="width: 10px;" />
                        <h3>:</h3>
                        <span style="width: 10px;" />
                        <h3>₹ {mainProduct.amount}</h3>
                    </div>
                    <div style="display: flex;">
                        <h3>Rating</h3>
                        <span style="width: 10px;" />
                        <h3>:</h3>
                        <span style="width: 10px;" />
                        <h3>{mainProduct.rating / 10}/5</h3>
                    </div>
                </div>
                <button on:click={Addtocart}>Add to cart</button>
            </div>
        </div>
        
    </div>
</div>


<style>
    .maincontainer {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        overflow-x: hidden;
        user-select: none;
    }

    .detialsContainer {
        margin-top: 20px;
        width: 60%;
        height: 100%;
        padding: 10px;
    }

    .card1 {
        width: 100%;
        height: 35%;
        display: flex;
    }
    .imagecon {
        min-width: 400px;
        max-width: 400px;
        width: 400px;
        flex: 1;
    }
    .imagecon img {
        width: 100%;
        height: 100%;
        object-fit: contain;
    }
    .detialscon {
        padding-left: 20px;
        flex: 2;
        display: flex;
        flex-direction: column;
        align-items: start;
        justify-content: space-around;
    }
    button {
        border-radius: 10px;
        background-color: white;
        border: 1px solid rgba(0, 0, 0, 0.5);
        font-size: 12px;
        font-weight: bold;
        padding: 12px 45px;
        letter-spacing: 1px;
        text-transform: uppercase;
        transition: transform 80ms ease-in;
    }
    button:active {
        background-color: azure;
    }
</style>
