import apiClient from "@/modules/api/apiClient";
import storage from "@/storage";

const getCategoryById = async (id) => {
  try {
    const res = await apiClient
      .get(`/private/categories/${id}`, {
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

const getCategoriesByQueryParams = async (page, limit) => {
  try {
    const res = await apiClient
      .get(`/private/categories?page=${page}&limit=${limit}`, {
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

const getCategories = async () => {
  try {
    const res = await apiClient
      .get(`/private/categories?page=1&limit=9999`, {
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

const postCategory = async (name) => {
  try {
    const res = await apiClient
      .post(
        `/private/categories`,
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

const patchCategory = async (id, name) => {
  try {
    const res = await apiClient
      .patch(
        `/private/categories/${id}`,
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

const deleteCategory = async (id) => {
  try {
    const res = await apiClient
      .delete(`/private/categories/${id}`, {
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

const apiCategories = {
  getCategoryById,
  getCategoriesByQueryParams,
  getCategories,
  postCategory,
  patchCategory,
  deleteCategory,
};

export default apiCategories;
