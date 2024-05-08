import { Form, Input } from "antd";

const RegisterPage = () => {
  return (
    <Form>
      <Form.Item
        name="username"
        rules={[{ required: true }, { whitespace: true }]}
      >
        <Input />
      </Form.Item>
      <Form.Item name="password" rules={[{ required: true }]}>
        <Input.Password />
      </Form.Item>
    </Form>
  );
};

export default RegisterPage;
