import useAuth from "@/hooks/useAuth";
import useAxiosPrivate from "@/hooks/useAxiosPrivate";
import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

const StravaCallback = () => {
  const axios = useAxiosPrivate();
  const [searchParams] = useSearchParams();
  let code = searchParams.get("code");
  let navigate = useNavigate();
  const { setAuth } = useAuth();

  useEffect(() => {
    if (!code) return;
    let controller = new AbortController();
    const oauth = async () => {
      try {
        await axios.get("/strava-oauth", {
          params: { code },
          signal: controller.signal,
        });
        navigate("/", { replace: true });

        setAuth((prev) => {
          return {
            ...prev,
            user: {
              ...prev.user,
              stravaLinked: true,
            },
          };
        });
      } catch (e) {
        console.log("e error:", e);
      }
    };

    oauth();

    return () => controller.abort();
  }, []);

  return <div>OAuthng Strava...</div>;
};

export default StravaCallback;
