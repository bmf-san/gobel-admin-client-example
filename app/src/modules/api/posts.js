import apiClient from "@/modules/api/apiClient";
import storage from "@/storage";

const getPostById = async (id) => {
  try {
    const res = await apiClient
      .get(`/private/posts/${id}`, {
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

const getPostsByQueryParams = async (page, limit) => {
  try {
    const res = await apiClient
      .get(`/private/posts?page=${page}&limit=${limit}`, {
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

const postPost = async (
  categoryId,
  tagIds,
  title,
  markdown,
  compiledMarkdown,
  status
) => {
  try {
    const res = await apiClient
      .post(
        `/private/posts`,
        {
          category_id: categoryId,
          tags: tagIds,
          title: title,
          md_body: markdown,
          html_body: compiledMarkdown,
          status: status,
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

const patchPost = async (
  id,
  categoryId,
  tagIds,
  title,
  markdown,
  compiledMarkdown,
  status
) => {
  try {
    const res = await apiClient
      .patch(
        `/private/posts/${id}`,
        {
          category_id: categoryId,
          tags: tagIds,
          title: title,
          md_body: markdown,
          html_body: compiledMarkdown,
          status: status,
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

const deletePost = async (id) => {
  try {
    const res = await apiClient
      .delete(`/private/posts/${id}`, {
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

const apiPosts = {
  getPostById,
  getPostsByQueryParams,
  postPost,
  patchPost,
  deletePost,
};

export default apiPosts;
