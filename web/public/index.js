axios.interceptors.request.use(config => {
  // Authorization Token
  const token = localStorage.getItem("user_token");

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});