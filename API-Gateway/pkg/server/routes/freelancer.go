package routes

import (
	handler "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/handlers"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Freelancer(
	api fiber.Router,
	profile *handler.ProfileHandler,
	project *handler.ProjectHandler,
	job *handler.JobsHandler,
) {
	profiles := api.Group("/profile")
	{
		profiles.Use(middlewares.AuthFreelancer)
		{
			profiles.Get("", profile.GetFreelancerProfile)
			profiles.Post("/education", profile.AddEducationDetails)
			profiles.Patch("/education", profile.UpdateEducation)
			profiles.Post("/experience", profile.AddExperience)
			profiles.Patch("/experience", profile.UpdateExperience)
			profiles.Delete("/experience", profile.RemoveExperience)
			profiles.Delete("/education", profile.DeleteEducation)
			profiles.Post("/description", profile.AddProfileDescription)
			profiles.Patch("/description", profile.EditProfileDescription)
			profiles.Put("/photo", profile.UpdateProfilePhoto)
			profiles.Post("/skill", profile.UpdateSkilltoProfile)
		}
	}

	jobs := api.Group("/job")
	{
		jobs.Post("accept-offer/:offer_id", middlewares.AuthFreelancer, job.AcceptOffer)
		jobs.Post("send-proposal", middlewares.AuthFreelancer, job.SendProposal)
		jobs.Get("/offers",middlewares.AuthFreelancer  ,job.GetAllJobOffersForFreelancer)
		jobs.Get("",job.GetJobs)
		jobs.Get("/:id",job.GetJob)
		jobs.Post("/invoice", middlewares.AuthFreelancer,job.SendInvoice)
	}

	contract:=api.Group("contracts")
	{
		contract.Use(middlewares.AuthFreelancer)
		{
			// contract.Get("",job.GetMyContractsForFreelancer)
		}
	}

	projects := api.Group("/project")
	{
		projects.Use(middlewares.AuthFreelancer)
		{
			projects.Post("",project.AddSingleProject)
			projects.Patch("/:id",project.EditProject)
			projects.Delete("/:id",project.RemoveProject)
			projects.Get("",project.ListProjects)
			projects.Get("/:id",project.ListProjectWithID)
			projects.Post("/buy/:id",project.BuyProject)
			// projects.Get("/payment/:id",project.ExecutePaymentForProject)
			projects.Get("/user/only",project.ListMyProjects)
		}
	}
	payment:=api.Group("/payment")
	payment.Post("/onboard-freelancers",profile.OnboardFreelancersToPaypal)
	api.Post("/payment/add",profile.AddPaymentEmailPaypal)

}
