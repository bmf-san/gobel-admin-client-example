const PUBLISH = "publish";
const DRAFT = "draft";

export default Object.freeze({
  STATUS_PUBLISH: PUBLISH,
  STATUS_PRIVATE: DRAFT,
  STATUSES: [PUBLISH, DRAFT],
});
