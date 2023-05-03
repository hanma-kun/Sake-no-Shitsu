<script lang="ts">
    import { GetCart } from "../../utils/requests/get";
    import { PostCart, RemoveCart } from "../../utils/requests/post";
    import type { Cart } from "../../utils/requests/types";
    let arr: Cart;
    let total: number = 0;
    function ChangeContent() {
        const elem = document.getElementById("left_value");
        elem.innerHTML=""
        arr.cart.forEach((value) => {
            const h3 = document.createElement("h4");
            const h4 = document.createElement("h4");
            const btn = document.createElement("button");
            const div = document.createElement("div");
            div.classList.add("cartbar");
            h3.innerText = value.product_name;
            h4.innerText = `₹ ${value.price}`;
            btn.innerHTML = `
<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 96 960 960" width="24"><path d="M304.615 896q-26.846 0-45.731-18.884Q240 858.231 240 831.385V336h-40v-40h160v-30.77h240V296h160v40h-40v495.385Q720 859 701.5 877.5 683 896 655.385 896h-350.77ZM680 336H280v495.385q0 10.769 6.923 17.692T304.615 856h350.77q9.23 0 16.923-7.692Q680 840.615 680 831.385V336ZM392.307 776h40.001V416h-40.001v360Zm135.385 0h40.001V416h-40.001v360ZM280 336v520-520Z"/></svg>`;
            btn.onclick = ()=>Activate(value._id)
            div.appendChild(h3);
            div.appendChild(h4);
            div.appendChild(btn);
            elem.appendChild(div);
        });
    }
    let payment_method:string
    async function Activate(value:string){
        await RemoveCart(value) 
        Init()
    }

    async function Init() {
        total=0
        arr = await GetCart();
        arr.cart.forEach((value) => {
            total += value.price;
        });
        ChangeContent();
    }
    Init();
    async function PlaceorderFunc(){
        await PostCart(arr.cart,total,payment_method)
    }

</script>

<!-- House   string             `bson:"house_name" json:"house_name"`
	Street  string             `bson:"street_name" json:"street_name"`
	City    string             `bson:"city_name" json:"city_name"`
	Pincode string             `bson:"pincode" json:"pincode"` -->

<div class="mainconttainer">
    <div class="container">
        <div class="left" >
            <h2 style="margin-bottom: 20px;text-align: center;">Your Cart</h2>
            <div id="left_value"></div>
            <div class="bottom">
                <h3>Total Amount :</h3>
                <h3>₹ {total}</h3>
            </div>
        </div>
        <div class="right">
            <div class="address" id="address">
                <!-- <h2 style="margin-bottom: 20px;text-align: center;">Your Detials</h2>
                <div class="newAddress">
                    <div><input type="text" placeholder="House Name"/></div>
                    <div><input type="text" placeholder="Street Name"/></div>
                    <div class="cstm"><input type="text" placeholder="City Name"/><input type="number" placeholder="Pincode"/></div>
                    <button>Add new Address </button>
                </div> -->
                <div class="payment">
                    <label for="">
                        <input type="radio" name="payment" value="COD" bind:group={payment_method}> COD
                    </label>
                    <label for="">
                        <input type="radio" name="payment" value="Digital" bind:group={payment_method}> Digital Payment
                    </label>
                </div>
            </div>

            <div class="bottom">
                <button on:click={PlaceorderFunc}>Place your Order</button>
            </div>
        </div>
    </div>
</div>

<style>
    .mainconttainer {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
    }
    .container {
        width: 90%;
        height: 100%;
        display: flex;
    }
    .left {
        flex: 1;
        padding: 10px;
        position: relative;
    }
    .left .bottom {
        position: absolute;
        width: 98%;
        height: 50px;
        bottom: 5px;
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-template-rows: 1fr;
        grid-column-gap: 30px;
        grid-row-gap: px;
        place-items: center;
    }
    .right {
        flex: 1;
        position: relative;
    }
    .right .bottom{
        position: absolute;
        width: 95%;
        height: 50px;
        bottom: 5px;
        left: 15px;
    }
    .right .bottom button{
        width: 100%;
        height: 100%;
        border-radius: 10px;
        background-color: white;
        border: 1px solid rgba(0, 0, 0, 0.5);
        font-size: 16px;
        font-weight: bold;
        padding: 10px;
        letter-spacing: 1px;
        text-transform: uppercase;
        transition: transform 80ms ease-in;
    }
    .right .bottom button:active{
        background-color: azure;
    }
    .payment{
        margin-top: 30px;
        display: flex;
        gap: 30px;
        justify-content: center;
    }
    /* .cstm{
        display: flex; 
        gap: 10px;
    }
    .newAddress  input {
        background-color: #fff;
        border: 1px solid rgba(0, 0, 0, 0.39);
        border-radius: 5px;
        padding: 12px 15px;
        margin: 8px 0;
        width: 100%;
    }
    .newAddress button{
        border-radius: 5px;
        background-color: white;
        border: 1px solid rgba(0, 0, 0, 0.5);
        font-size: 14px;
        font-weight: bold;
        padding: 5px;
        letter-spacing: 1px;
        text-transform: uppercase;
        transition: transform 80ms ease-in;
    }
    .newAddress button:active{
        background-color: azure;
    } */
</style>
