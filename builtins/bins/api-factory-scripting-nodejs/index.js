"use strict";

// const vm2 = require("vm2");
const grpc = require("grpc");
const messages = require("./proto/scripting_pb");
const services = require("./proto/scripting_grpc_pb");

const run = () => {};

const getServer = bindAddress => {
  const server = new grpc.Server();
  server.addService(services.ScriptingProviderService, {
    run
  });
  server.bind(bindAddress, grpc.ServerCredentials.createInsecure());
  return server;
};

const server = plugin.getServer("localhost:50051");
server.start();
