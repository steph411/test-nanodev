"use strict";

/**
 * expertise-area service
 */

const { createCoreService } = require("@strapi/strapi").factories;
const axios = require("axios");

const SERVICE_URL = `${process.env.INTERNAL_API_URL}/areas`;

module.exports = createCoreService(
  "api::expertise-area.expertise-area",
  ({ strapi }) => ({
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
    async update(...args) {
      console.log({ args });
    },
    async create(params) {
      console.log({ params });
      const response = await axios.post(SERVICE_URL, params.data);
      console.log({ response });
      if (response.status === 201) {
        return response.data;
      }
    },
  })
);
