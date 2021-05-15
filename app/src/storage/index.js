const storage = {
  setIsSignin(bool) {
    localStorage.setItem("is_signin", bool ? 1 : 0);
  },
  getIsSignin() {
    return Boolean(parseInt(localStorage.getItem("is_signin"), 10));
  },
  setAccessToken(token) {
    localStorage.setItem("access_token", token);
  },
  getAccessToken() {
    return localStorage.getItem("access_token");
  },
  setRefreshToken(token) {
    localStorage.setItem("refresh_token", token);
  },
  getRefreshToken() {
    return localStorage.getItem("refresh_token");
  },
  removeAccessToken() {
    localStorage.removeItem("access_token");
  },
  removeRefreshToken() {
    localStorage.removeItem("refresh_token");
  }
};

export default storage;
