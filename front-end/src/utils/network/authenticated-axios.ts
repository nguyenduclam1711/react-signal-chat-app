import axios, { AxiosError } from "axios";
import { defaultConfigForAxios, handleRefreshTokenError, handleRefreshTokenSuccess, refreshAccessToken } from "./helpers";

export const AuthenticatedAxios = axios.create({
  ...defaultConfigForAxios,
});

AuthenticatedAxios.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error: AxiosError) => {
    if (error.response && error.response.status === 401) {
      try {
        // TODO: get refresh token from header request cookie, can't get from js because cookie config http only
        const refreshToken = "";
        const response = await refreshAccessToken(refreshToken);

        // set new refresh token and access token to cookie and remember to set http only
        handleRefreshTokenSuccess(response);

        // retry the original request
        const originalRequest = error.config;
        if (originalRequest) {
          // remove both refresh token and access token from cookie
          handleRefreshTokenError();
          return AuthenticatedAxios(originalRequest);
        }
        return Promise.reject(error);
      } catch (refreshError) {
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  },
);
