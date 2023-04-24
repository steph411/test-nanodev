"use strict";

/**
 * request service
 */

const { createCoreService } = require("@strapi/strapi").factories;
const axios = require("axios");
const SERVICE_URL = `${process.env.INTERNAL_API_URL}/members`;

module.exports = createCoreService("api::request.request", ({ strapi }) => ({
  async find(...args) {
    console.log({ args });
    const response = await axios.get(SERVICE_URL, {
      method: "GET",
      mode: "cors",
    });
    console.log({ response });
    if (response.status === 200) {
      return { results: response.data.data };
    }
  },
  async findOne(...args) {
    console.log({ args });
  },
  async update(entityId, params) {
    console.log({ args });
    const response = await axios.post(
      `${SERVICE_URL}/${entityId}`,
      params.data
    );
    console.log({ response });
    if (response.status === 201) {
      return response.data;
    }
  },
  async create(params) {
    console.log({ params });
    const response = await axios.post(SERVICE_URL, params.data);
    console.log({ response });
    if (response.status === 201) {
      return response.data;
    }
  },
}));
