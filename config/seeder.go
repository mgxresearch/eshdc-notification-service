package config

import (
	"fmt"

	"github.com/eshdc/notification-service/models"
)

func SeedTemplates() {
	fmt.Println("🌱 Seeding Notification Templates...")

	templates := []models.NotificationTemplate{
		{
			Name:    "welcome_email",
			Subject: "Welcome to ESHDC Digitization Portal",
			Body:    "<h1>Welcome {{user_name}}!</h1><p>Your account has been successfully created. You can now apply for properties digitally.</p>",
			Type:    "email",
		},
		{
			Name:    "application_received",
			Subject: "Your ESHDC Property Application: {{estate_name}}",
			Body:    "<h1>Hello {{user_name}},</h1><p>We have received your application for a {{plot_type}} plot in {{estate_name}}. Our team is currently reviewing your KYC documents.</p>",
			Type:    "email",
		},
		{
			Name:    "payment_confirmed",
			Subject: "Payment Confirmed for {{estate_name}}",
			Body:    "<h1>Payment Successful!</h1><p>Your payment of {{amount}} for {{estate_name}} has been confirmed. You will be notified once your plot allocation is finalized.</p>",
			Type:    "email",
		},
		{
			Name:    "incoming_transfer",
			Subject: "Incoming Property Transfer: {{property_name}}",
			Body:    "<h1>Property Transfer Initiated</h1><p>A transfer of {{property_name}} to your account has been initiated. This request is currently undergoing administrative review and due diligence. You will be notified once the finalization is complete.</p>",
			Type:    "email",
		},
		{
			Name:    "transfer_approved",
			Subject: "Property Transfer Finalized: {{property_name}}",
			Body:    "<h1>Transfer Completed!</h1><p>Your property transfer for {{property_name}} has been approved by ESHDC Admin. You are now the official owner.</p>",
			Type:    "email",
		},
		{
			Name:    "internal_memo",
			Subject: "ESHDC Internal Dispatch: {{subject}}",
			Type:    "email",
			Body: `<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: 'Inter', Arial, sans-serif; background-color: #f8fafc; margin: 0; padding: 0; }
  .container { max-width: 600px; margin: 40px auto; background: white; border-radius: 16px; overflow: hidden; border: 1px solid #e2e8f0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); }
  .header { background-color: #065f46; color: white; padding: 32px; text-align: center; }
  .logo { font-weight: 900; font-size: 16px; letter-spacing: -0.025em; text-transform: uppercase; margin-bottom: 8px; }
  .header-sub { font-size: 10px; font-weight: 700; opacity: 0.8; letter-spacing: 0.2em; text-transform: uppercase; }
  .content { padding: 40px; color: #1e293b; line-height: 1.6; }
  .meta { background-color: #f1f5f9; padding: 20px; border-radius: 12px; margin-bottom: 32px; font-size: 12px; }
  .meta-row { display: flex; margin-bottom: 8px; }
  .meta-label { width: 100px; color: #64748b; font-weight: 800; text-transform: uppercase; }
  .meta-value { font-weight: 700; color: #0f172a; }
  .body-content { font-size: 15px; border-top: 2px solid #f1f5f9; padding-top: 32px; margin-top: 32px; }
  .footer { padding: 32px; background-color: #f8fafc; border-top: 1px solid #e2e8f0; text-align: center; font-size: 11px; color: #64748b; font-weight: 600; text-transform: uppercase; }
  .btn { display: inline-block; padding: 14px 28px; background-color: #059669; color: white; text-decoration: none; border-radius: 10px; font-weight: 800; font-size: 12px; text-transform: uppercase; letter-spacing: 0.05em; margin-top: 24px; text-align: center; }
</style>
</head>
<body>
  <div class="container">
    <div class="header">
      <div class="logo">Enugu State Housing Development Corporation</div>
      <div class="header-sub">Staff Command Center - Internal Dispatch</div>
    </div>
    <div class="content">
      <h2 style="font-weight: 900; font-size: 22px; margin-top: 0; text-transform: uppercase; letter-spacing: -0.01em; color: #0f172a;">{{subject}}</h2>
      <div class="meta">
        <div style="display: table; width: 100%;">
          <div style="display: table-row;"><div style="display: table-cell; width: 100px; padding-bottom: 8px; color: #64748b; font-weight: 800; text-transform: uppercase; font-size: 10px;">Category:</div><div style="display: table-cell; font-weight: 700; color: #0f172a;">{{category}}</div></div>
          <div style="display: table-row;"><div style="display: table-cell; width: 100px; padding-bottom: 8px; color: #64748b; font-weight: 800; text-transform: uppercase; font-size: 10px;">Serial:</div><div style="display: table-cell; font-weight: 700; color: #0f172a;">{{memo_serial}}</div></div>
          <div style="display: table-row;"><div style="display: table-cell; width: 100px; color: #64748b; font-weight: 800; text-transform: uppercase; font-size: 10px;">From:</div><div style="display: table-cell; font-weight: 700; color: #0f172a;">{{sender_name}}</div></div>
        </div>
      </div>
      <div class="body-content">
        {{content}}
      </div>
      <center><a href="https://portal.enugustate.gov.ng/admin" class="btn">View in Command Center</a></center>
    </div>
    <div class="footer">
      Official Internal Correspondence<br/>
      No. 21 Aguleri St, Independence Layout, Enugu<br/>
      &copy; 2024 ESHDC. Restricted Access.
    </div>
  </div>
</body>
</html>`,
		},
		{
			Name:    "external_memo",
			Subject: "Official Correspondence from ESHDC: {{subject}}",
			Type:    "email",
			Body: `<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: 'Inter', Arial, sans-serif; background-color: #f8fafc; margin: 0; padding: 0; }
  .container { max-width: 600px; margin: 40px auto; background: white; border-radius: 16px; overflow: hidden; border: 1px solid #e2e8f0; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); }
  .header { background-color: #065f46; color: white; padding: 32px; text-align: center; }
  .logo { font-weight: 900; font-size: 16px; letter-spacing: -0.025em; text-transform: uppercase; margin-bottom: 8px; }
  .header-sub { font-size: 10px; font-weight: 700; opacity: 0.8; letter-spacing: 0.2em; text-transform: uppercase; }
  .content { padding: 40px; color: #1e293b; line-height: 1.6; }
  .meta { background-color: #f1f5f9; padding: 20px; border-radius: 12px; margin-bottom: 32px; font-size: 12px; }
  .body-content { font-size: 15px; border-top: 2px solid #f1f5f9; padding-top: 32px; margin-top: 32px; }
  .footer { padding: 32px; background-color: #f8fafc; border-top: 1px solid #e2e8f0; text-align: center; font-size: 11px; color: #64748b; font-weight: 600; text-transform: uppercase; }
  .btn { display: inline-block; padding: 14px 28px; background-color: #059669; color: white; text-decoration: none; border-radius: 10px; font-weight: 800; font-size: 12px; text-transform: uppercase; letter-spacing: 0.05em; margin-top: 24px; text-align: center; }
</style>
</head>
<body>
  <div class="container">
    <div class="header">
      <div class="logo">Enugu State Housing Development Corporation</div>
      <div class="header-sub">Official Citizen Correspondence</div>
    </div>
    <div class="content">
      <h2 style="font-weight: 900; font-size: 22px; margin-top: 0; text-transform: uppercase; letter-spacing: -0.01em; color: #0f172a;">{{subject}}</h2>
      <div class="meta">
        <div style="display: table; width: 100%;">
          <div style="display: table-row;"><div style="display: table-cell; width: 100px; padding-bottom: 8px; color: #64748b; font-weight: 800; text-transform: uppercase; font-size: 10px;">Reference:</div><div style="display: table-cell; font-weight: 700; color: #0f172a;">{{memo_serial}}</div></div>
          <div style="display: table-row;"><div style="display: table-cell; width: 100px; color: #64748b; font-weight: 800; text-transform: uppercase; font-size: 10px;">Office:</div><div style="display: table-cell; font-weight: 700; color: #0f172a;">{{sender_name}}</div></div>
        </div>
      </div>
      <div class="body-content">
        {{content}}
      </div>
      <center><a href="https://shdcadmin.enugustate.gov.ng/" class="btn">Login to Citizen Portal</a></center>
    </div>
    <div class="footer">
      Enugu State Housing Development Corporation (ESHDC)<br/>
      No. 21 Aguleri St, Independence Layout, Enugu<br/>
      &copy; 2024 ESHDC. Serving the People.
    </div>
  </div>
</body>
</html>`,
		},
		{
			Name:    "mfa_otp",
			Subject: "Your ESHDC Verification Code: {{otp}}",
			Body:    "<h1>Secure Access Verification</h1><p>Hello,</p><p>Your verification code is: <strong>{{otp}}</strong></p><p>This code will expire in 10 minutes. If you did not request this, please secure your account immediately.</p>",
			Type:    "email",
		},
		{
			Name:    "password_reset",
			Subject: "Reset Your ESHDC Password",
			Body:    "<h1>Password Reset Request</h1><p>Hello,</p><p>We received a request to reset your password. Use the code below to proceed:</p><p><strong>{{otp}}</strong></p><p>If you did not request a password reset, you can safely ignore this email.</p>",
			Type:    "email",
		},
		{
			Name:    "digitization_complete",
			Subject: "Scan Successful",
			Body:    "File {{file_number}} has been successfully digitized and added to the vault.",
			Type:    "in_app",
		},
		{
			Name:    "approval_required",
			Subject: "Approval Required",
			Body:    "Admin needs to approve the plot update for {{plot_id}}.",
			Type:    "in_app",
		},
	}

	for _, t := range templates {
		var existing models.NotificationTemplate
		if err := DB.Where("name = ?", t.Name).First(&existing).Error; err != nil {
			DB.Create(&t)
		} else {
			// Update existing to ensure Type is set
			DB.Model(&existing).Update("type", t.Type)
		}
	}

	fmt.Println("✅ Notification Seeding Completed!")
}
