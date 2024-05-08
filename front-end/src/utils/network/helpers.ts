import { AxiosRequestConfig } from "axios";
import cookie from "cookie";
import dayjs from "dayjs";

export const defaultConfigForAxios: AxiosRequestConfig = {
  baseURL: process.env.API_BASE_URL,
};

export const refreshAccessToken = async (refreshToken: string): Promise<{ accessToken: string, refreshToken: string }> => {
  // TODO: refresh access token here
  return {
    accessToken: "",
    refreshToken: "",
  };
};

export const handleRefreshTokenSuccess = (param: {
  accessToken: string,
  refreshToken: string,
}) => {
  cookie.serialize("accessToken", param.accessToken, {
    httpOnly: true,
  });
  cookie.serialize("refreshToken", param.refreshToken, {
    httpOnly: true,
  });
};

export const handleRefreshTokenError = () => {
  const yesterday = dayjs().subtract(1, "days").toDate();

  cookie.serialize("accessToken", "", {
    httpOnly: true,
    expires: yesterday,
  });
  cookie.serialize("refreshToken", "", {
    httpOnly: true,
    expires: yesterday,
  });
};
