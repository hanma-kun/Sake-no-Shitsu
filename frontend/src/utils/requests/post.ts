import { navigate } from "svelte-navigator"
import { token, Alerts } from "../../store"
import httpClient from "./client"
import type { LoginCredType, LoginResponse, ProductType, SigninCredType } from "./types"



export async function LoginUser(Logindata: LoginCredType) {
    try {
        const response = await httpClient.post("/login", JSON.stringify(Logindata))
        if (response.status == 200) {
            const data = (await response.data) as LoginResponse
            localStorage.setItem("token", data.access_token)
            localStorage.setItem("refresh_token", data.refresh_token)
            token.set(data.access_token)
        }
    } catch (e) {
        Alerts.update((value) => {
            return { ...value, title: "Error Occured", message: e.response.data.error, show: true }
        })
    }
}

export async function SigninUser(SigninCreds: SigninCredType) {
    try {
        const response = await httpClient.post("/signup", JSON.stringify(SigninCreds))
        if (response.status == 200) {
            const data = (await response.data) as LoginResponse
            localStorage.setItem("token", data.access_token)
            localStorage.setItem("refresh_token", data.refresh_token)
            token.set(data.access_token)
        }
    } catch (e) {
        Alerts.update((value) => {
            return { ...value, title: "Error Occured", message: e.response.data.error, show: true }
        })
    }
}

export async function AddtoCart(id: string) {
    try {
        const response = await httpClient.post(`/cart_add?id=${id}`)

        if (response.status == 200) {
            const data = await response.data
            Alerts.update((value) => {
                return { ...value, title: "Added to Cart", message: `${data[0].product_name} has beem added to your cart successfully `,show:true }
            })

            return true
        }
        return false
    } catch (e) {
        return false
    }
}

export async function RemoveCart (value:string) {
    try{
        const response = await httpClient.post(`/cart_remove?id=${value}`)
        if (response.status == 200){
            return true
        }
        return false
    }catch(e){
        console.log(e)
    }
 

    
}

export async function PostCart(cart:ProductType[],total:number,payment:string){
    try{
        
        const response = await httpClient.post('/cart_order',{
            order_cart:cart,
            price:total,
            payment:payment})
        if (response.status == 200){
            Alerts.update(value=>{
                return {...value,title:"Order Places successfully",message:`Your order of rupees ${total} is placed Successfully`,show:true,handleOk:()=>{
                    navigate("/",{replace:true})
                }}
            })
        }
        } catch(e){
        console.log(e)
    }
}