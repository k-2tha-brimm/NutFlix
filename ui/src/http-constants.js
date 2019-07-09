import axios from 'axios';

let baseURL = "http://localhost:5000";

export const HTTP = axios.create(
    {
        baseURL: baseURL
    });