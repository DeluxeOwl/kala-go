// @ts-nocheck
import { BACKEND_URL } from "../url";

const URLS = [
  `${BACKEND_URL}/typeconfig/batch`,
  // `${BACKEND_URL}/subject/batch`,
  //   `${BACKEND_URL}/tuple/batch`,
  //   `${BACKEND_URL}/permission-check/batch`,
];

const postReq = (url, body) =>
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

const fetchAll = async (data) => {
  for (const [i, url] of URLS.entries()) {
    const res = await postReq(url, data[i]);
    const body = await res.json();

    if (res.status >= 400) {
      throw new Error(body.message);
    }
  }
};

export { fetchAll };
