import httpClient from "./client"
import type { Cart, LoginResponse, ProductType } from "./types"
import { Alerts, token } from "../../store"
let key;
token.subscribe((value) => {
    key = value;
})
export async function VerifyToken() {
    if (key) {
        try {
            const _ = await httpClient.get("/verifytoken")
        } catch (e) {
            if (e.response.status == 401) {
                try {
                    const refresh_token = localStorage.getItem("refresh_token")
                    const refresponse = await httpClient.post("/refreshtoken", {
                        refresh_token,
                    })
                    const data = (await refresponse.data) as LoginResponse
                    localStorage.setItem("token", data.access_token)
                    localStorage.setItem("refresh_token", data.refresh_token)
                    token.set(data.access_token)
                } catch (e) {
                    localStorage.clear()
                }
            }
        }
    }

}

export async function GetProduct(category: string) {
    try {
        const response = await httpClient.get(`/products?category=${category}`)
        if (response.status == 200) {
            //@ts-ignore
            const products: ProductType[] = response.data.map((product: any) => {
                return {
                    _id: product._id,
                    category: product.category,
                    image: product.image,
                    price: product.price,
                    product_name: product.product_name,
                    rating: product.rating
                };
            });
            return products
        }
        return []
    } catch (e) {
        console.log(e)
    }

}

export async function Getdetials(value:string) {
    try{
        let kk:ProductType
        const response = await httpClient.get(`productdetials?id=${value}`)
        if (response.status == 200){
            const data = (await response.data) as ProductType
            return data
        }
        return kk
    }catch(e){}
}

export async function GetCart() {
    try{
        let data:Cart
        const response = await httpClient.get("/cart_all")
        if (response.status == 200 ){
            data = (await response.data) as Cart
        }
        return data        
    }catch(e){
        Alerts.update((value)=>{
            return {...value,title:"Error",message:"We are currently unable to update the cart",show:true}
        })
    }
    
}