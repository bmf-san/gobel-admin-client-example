import apiClient from "@/modules/api/apiClient";
import storage from "@/storage";

const getCommentById = async id => {
  try {
    const res = await apiClient
      .get(`/private/comments/${id}`, {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken()
        }
      })
      .then(res => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const getCommentsByQueryParams = async (page, limit) => {
  try {
    const res = await apiClient
      .get(`/private/comments?page=${page}&limit=${limit}`, {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken()
        }
      })
      .then(res => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const patchComment = async (id, status) => {
  try {
    const res = await apiClient
      .patch(
        `/private/comments/${id}/status`,
        {
          status: status
        },
        {
          headers: {
            Authorization: "Bearer " + storage.getAccessToken()
          }
        }
      )
      .then(res => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const apiComments = {
  getCommentById,
  getCommentsByQueryParams,
  patchComment
};

export default apiComments;
