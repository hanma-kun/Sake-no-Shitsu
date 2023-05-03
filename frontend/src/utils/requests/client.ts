import axios from "axios";
import { token } from "../../store";
import type {AxiosInstance,AxiosRequestConfig} from "axios";


let jwt:string;
token.subscribe((value)=>{
    jwt=value;
})

function getJwt(){
    return jwt
}

class HttpClient{
    private readonly client:AxiosInstance;

    constructor(){
        this.client = axios.create({
            baseURL:"http://localhost:3000", // need to change 
            headers:{
                "Content-Type":"application/json"
            }
        })
    }

    #setHeader(){
        this.client.defaults.headers.common={
            ...this.client.defaults.headers.common,
            Authorization: getJwt()
        }
    }

    get<T>(url:string,config?:AxiosRequestConfig){
        this.#setHeader()
        return this.client.get<T>(url,config);
    }

    post<T>(url:string,data?:any,config?:AxiosRequestConfig){
        this.#setHeader()
        return this.client.post<T>(url,data,config);
    }
}

export default new HttpClient();