import { writable } from "svelte/store";

type Alertsettings = {
    title: string;
    message: string;
    handleOk: () => void;
    handleCancel: () => void;
    show: boolean;
}


export const token = writable(localStorage.getItem('token'))
export const Alerts = writable<Alertsettings>({ title: "", message: "", show: false, handleCancel: () => { }, handleOk: () => { } })
