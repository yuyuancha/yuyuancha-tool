import axios from "axios";

const dogHeroApi = axios.create({
    baseURL: `${import.meta.env.VITE_BACKEND_API}/v1/dogHero`,
    headers: {
        "Content-Type": "application/json; charset=utf-8",
        Accept: "application/json",
    },
});

export default dogHeroApi;
