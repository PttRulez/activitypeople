import { RegisterData } from "@/validation/auth";
import { useForm } from "react-hook-form";

const Register = () => {
  const {
    formState: { errors },
    handleSubmit,
    watch,
  } = useForm<RegisterData>();

  return (
    <div className='flex justify-center mt-[calc(100vh-100vh+8rem)]'>
      <div className='max-w-screen-sm w-full bg-base-300 py-10 px-16 rounded-xl'>
        <h1 className='text-center text-primary text-3xl mb-10'>Register</h1>

        <form hx-post='/login' className='space-y-5'>
          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Имя</span>
            </div>
            <input
              type='text'
              name='name'
              value={watch("name")}
              placeholder='ваше имя/ник на сайте'
              className='input input-bordered w-full'
              autoComplete='off'
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.email?.message}
              </span>
            </div>
          </label>

          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Email</span>
            </div>
            <input
              type='email'
              name='email'
              value={watch("email")}
              placeholder='type your email'
              className='input input-bordered w-full'
              autoComplete='off'
              required
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.email?.message}
              </span>
            </div>
          </label>

          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Пароль</span>
            </div>
            <input
              type='password'
              name='password'
              value={watch("password")}
              placeholder='type your password'
              className='input input-bordered w-full'
              autoComplete='off'
              required
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.password?.message}
              </span>
            </div>
          </label>

          <label className='form-control w-full'>
            <div className='label'>
              <span className='label-text'>Подвтерждение пароля</span>
            </div>
            <input
              type='password'
              name='confirmPassword'
              placeholder='confirm your password'
              className='input input-bordered w-full'
              autoComplete='off'
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.confirmPassword?.message}
              </span>
            </div>
          </label>

          <button type='submit' className='btn btn-primary w-full'>
            register <i className='fa-solid fa-arrow-right'></i>
          </button>
          <div className='divider'>OR</div>
          <a href='/login' className='btn btn-outline w-full'>
            Login
          </a>
        </form>
      </div>
    </div>
  );
};

export default Register;
