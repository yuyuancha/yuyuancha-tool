import axios from "axios";

export const dogHeroApi = axios.create({
    baseURL: `${import.meta.env.VITE_BACKEND_API}/v1/dogHero`,
    headers: {
        "Content-Type": "application/json; charset=utf-8",
        Accept: "application/json",
    },
});

export const govApi = axios.create({
    baseURL: `${import.meta.env.VITE_BACKEND_API}/v1/gov`,
    headers: {
        "Content-Type": "application/json; charset=utf-8",
        Accept: "application/json",
    },
});
