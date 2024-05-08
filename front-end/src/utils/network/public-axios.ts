import axios from "axios";
import { defaultConfigForAxios } from "./helpers";

export const PublicAxios = axios.create(defaultConfigForAxios);
