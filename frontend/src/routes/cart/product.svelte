<script lang="ts">
    import { useLocation,useNavigate } from "svelte-navigator";
    import { GetProduct } from "../../utils/requests/get";
    import type { ProductType } from "../../utils/requests/types";
    import Loader from "../../utils/loader.svelte";
    let location = useLocation();
    let navigate = useNavigate()
    let data: ProductType[] = [];
    let loading = false;
    location.subscribe(async (value) => {
        loading = true;
        data = await GetProduct(value.search.replace("?", ""));
        loading = false;
    });

    function Preview(value:string){
        navigate("/preview",{
            state:{data:value},
            replace:true,
        })
    }
</script>

{#if loading}
    <Loader />
{/if}
<div class="container">
    <div class="maincontainer">
        {#each data as product}
            <div class="card" on:click={()=>Preview(product._id)} on:keypress>
                <div class="imagecontainer">
                    <img src={product.image} alt={product.product_name} draggable="false" />
                </div>
                <div class="textcontainer">
                    <div class="header">
                        <h3>{product.product_name}</h3>
                    </div>
                    <div class="detials">
                        <div>
                            <h4>Price : ₹ {product.price}</h4>
                        </div>
                        <div>
                            <h4>Rating : {product.rating / 10}/5</h4>
                        </div>
                    </div>
                </div>
            </div>
        {/each}
    </div>
</div>

<style>
    .container {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        overflow-x: hidden;
    }
    .maincontainer {
        margin-top: 20px;
        width: 60%;
        padding: 10px;
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-column-gap: 20px;
        grid-row-gap: 20px;
    }
    .card {
        width: 360px;
        height: 400px;
        background-color: white;
        border-radius: 10px;
        border: 1px solid rgba(0, 0, 0, 0.253);
        box-shadow: 0 10px 20px rgba(0, 0, 0, 0.05),
            0 10px 10px rgba(0, 0, 0, 0.05);
        padding: 10px;
        position: relative;
        user-select: none;
    }
    
    .card:hover {
        border-color: #ffd700 !important;
    }
    .card:active{
        border: none;
    }

    .imagecontainer {
        width: 320px;
        height: 250px;
    }
    .imagecontainer img {
        width: 100%;
        height: 100%;
        object-fit: contain;
        overflow: hidden;
    }
    .textcontainer {
        width: 320px;
        position: absolute;
        padding: 10px 20px;
        bottom: 0;
    }
    .header {
        display: flex;
        justify-content: center;
        margin-bottom: 20px;
    }
    .detials {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 5px;
    }
</style>
