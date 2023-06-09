import apiClient from "@/modules/api/apiClient";
import storage from "@/storage";

const postSignin = async (email, password) => {
  try {
    const res = await apiClient
      .post(
        `/signin`,
        {
          email: email,
          password: password,
        },
        {
          headers: {
            Authorization: "Bearer " + storage.getAccessToken(),
          },
        }
      )
      .then((res) => {
        storage.setAccessToken(res.data.access_token);
        storage.setRefreshToken(res.data.refresh_token);
        storage.setIsSignin(true);
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const apiSignin = {
  postSignin,
};

export default apiSignin;
