import resolve from "@rollup/plugin-node-resolve";
import babel from "@rollup/plugin-babel";
import commonjs from "@rollup/plugin-commonjs";
import json from "@rollup/plugin-json";
import replace from "@rollup/plugin-replace";

export default {
  input: "src/index.js",
  output: {
    file: "public/bundle.js",
    format: "cjs",
  },
  plugins: [
    replace({
      preventAssignment: true,
      "process.env.NODE_ENV": JSON.stringify("production"),
    }),
    resolve({ jsnext: true, preferBuiltins: true, browser: true }),
    json(),
    babel({ babelHelpers: "runtime", exclude: "**/node_modules/**" }),
    commonjs(),
  ],
};
