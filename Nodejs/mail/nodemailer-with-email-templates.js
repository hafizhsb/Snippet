const nodemailer = require("nodemailer");
const Email = require('email-templates');

let transporter = nodemailer.createTransport({
  host: 'smtp.gmail.com',
  port: 465,
  secure: true,
  auth: {
    user: process.env.EMAIL,
    pass: process.env.EMAIL_PASSWORD
  },
})

const email = new Email({
  template: 'test',
  message: {
    'from': 'Test'
  },
  transport: transporter,
  send: true,
})

module.exports = {
  sendMail(to, subject, content, otherConfig = {}) {
    return new Promise(resolve => {
      email.send({
        message: {
          to,
          subject,
        },
        locals: {
          name: content.name,
          url: content.url,
        }
      })
      .then(console.log)
      .catch(console.error);
    });
  },
};
