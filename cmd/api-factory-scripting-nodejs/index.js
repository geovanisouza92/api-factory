"use strict";

// const vm2 = require("vm2");
const grpc = require("grpc");
const messages = {
  health: require("./proto/health_pb"),
  scripting: require("./proto/scripting_pb")
};
const services = {
  health: require("./proto/health_grpc_pb"),
  scripting: require("./proto/scripting_grpc_pb")
};

const getServer = bindAddress => {
  const server = new grpc.Server();
  server.addService(services.health.HealthService, {
    check(call, callback) {
      const reply = new messages.health.HealthCheckResponse();
      reply.setStatus(
        messages.health.HealthCheckResponse.ServingStatus.SERVING
      );
      callback(null, reply);
    }
  });
  server.addService(services.scripting.ScriptingProviderService, {
    run(call, callback) {}
  });
  server.bind(bindAddress, grpc.ServerCredentials.createInsecure());
  return server;
};

const NETWORK_ADDR = process.env.NETWORK_ADDR || "127.0.0.1:50051";

const server = getServer(NETWORK_ADDR);
server.start();
console.log(`1|${process.env.APP_PROTOCOL_VERSION}|tcp|${NETWORK_ADDR}|grpc`);