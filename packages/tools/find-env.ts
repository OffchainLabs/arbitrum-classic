import findUp = require("find-up");
export const findEnv = () => findUp.sync(process.env.ENV_FILE || '.env');