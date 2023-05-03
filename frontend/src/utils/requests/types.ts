export type LoginCredType = {
    email: string,
    password: string
}

export type LoginResponse = {
    session_id: string,
    user_id: string,
    access_token: string,
    access_token_exp: string,
    refresh_token: string,
    refresh_token_exp: string
}

export type SigninCredType = {
    firstname: string,
    lastname: string,
    email: string,
    phone: string,
    password: string
}

export type ProductType = {
    _id: string,
    category: string,
    image: string,
    price: number,
    product_name: string,
    rating: number
}



export type AddressType = {
    _id: string,
    house_name: string,
    street_name: string,
    city_name: string,
    pincode: string
}
export type Cart = {
    cart: ProductType[]
    address: AddressType[]
}