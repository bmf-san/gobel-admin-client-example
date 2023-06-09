import apiClient from "@/modules/api/apiClient";
import storage from "@/storage";

const getTagById = async (id) => {
  try {
    const res = await apiClient
      .get(`/private/tags/${id}`, {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken(),
        },
      })
      .then((res) => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const getTagsByQueryParams = async (page, limit) => {
  try {
    const res = await apiClient
      .get(`/private/tags?page=${page}&limit=${limit}`, {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken(),
        },
      })
      .then((res) => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const getTags = async () => {
  try {
    const res = await apiClient
      .get(`/private/tags?page=1&limit=9999`, {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken(),
        },
      })
      .then((res) => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const postTag = async (name) => {
  try {
    const res = await apiClient
      .post(
        `/private/tags`,
        {
          name: name,
        },
        {
          headers: {
            Authorization: "Bearer " + storage.getAccessToken(),
          },
        }
      )
      .then((res) => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const patchTag = async (id, name) => {
  try {
    const res = await apiClient
      .patch(
        `/private/tags/${id}`,
        {
          name: name,
        },
        {
          headers: {
            Authorization: "Bearer " + storage.getAccessToken(),
          },
        }
      )
      .then((res) => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const deleteTag = async (id) => {
  try {
    const res = await apiClient
      .delete(`/private/tags/${id}`, {
        headers: {
          Authorization: "Bearer " + storage.getAccessToken(),
        },
      })
      .then((res) => {
        return res;
      });
    return res;
  } catch (e) {
    throw new Error(e);
  }
};

const apiTags = {
  getTagById,
  getTagsByQueryParams,
  getTags,
  postTag,
  patchTag,
  deleteTag,
};

export default apiTags;
