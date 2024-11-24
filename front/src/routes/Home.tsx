import useAuth from "@/hooks/useAuth";
import { Link, Navigate } from "react-router-dom";

let user = {
  strava: {
    accessToken: false,
  },
  name: "Aleksandr",
};

const stravaOAuthLink = process.env.REACT_APP_STRAVA_OAUTH_LINK;

const Home = () => {
  const { auth } = useAuth();
  return user.strava.accessToken ? (
    <div className='text-center'>
      <h1 className='text-2xl mt-10 mb-10'>Приветствуем вас, {user.name}</h1>
      <a
        href='/activities'
        hx-target='body'
        hx-push-url='true'
        className='btn btn-primary'
      >
        Ваши активити
      </a>
    </div>
  ) : (
    <div className='text-center'>
      {!auth.user.stravaLinked ? (
        <>
          <h1 className='text-2xl mt-10 mb-10'>
            Здарова, атлет. Чтобы посмотреть свои активности, необходимо
            законнектить сраву
          </h1>

          <a href={stravaOAuthLink} className='btn btn-primary'>
            Привяжите свой аккаунт Strava
          </a>
        </>
      ) : (
        <Navigate to='/activities' />
      )}
    </div>
  );
};

export default Home;
