provider "aws" {
  region              = local.aws_region
  allowed_account_ids = local.aws_allowed_account_ids

  assume_role {
    role_arn = local.aws_assume_role
  }
}

locals {
  secret_name             = "terraform/production/ap/documents-studio-fetch"
  aws_region              = "ap-northeast-1"
  aws_assume_role         = "arn:aws:iam::862962005596:role/tf-terraform-developer"
  aws_allowed_account_ids = ["862962005596"]

  application_environment = {
    "COMMON_CIDR"                                                           = "10.0.0.0/8"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__AmazonProfileName"         = "noProfile"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__Bucket"                    = "finalcloud-production-ap"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__CleanUPInterval"           = "3600000"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__ConcurrentServiceRequests" = "100"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__DocGenJobTimeOut"          = "1800000"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__DownloadTimeout"           = "180000"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__EnableProfileConfig"       = "false"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__FetchJobTimeOut"           = "1800000"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__LocalPath"                 = "c:\\\\Projects\\\\GeneratedFiles\\\\"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__PreAuthenticated"          = "false"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__Region"                    = "us-east-1"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__RetryUploadDownload"       = "3"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__Storage"                   = "0"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__ThreadCount"               = 4
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__UploadTimeout"             = "600000"
    "COMMON_DEFAULT_ProfileConfig__AmazonConfig__VerifyFile"                = "true"
    "FETCH_Jwt__Issuer"                                                     = "DocumentsStudioScheduler"
  }
}

data "aws_secretsmanager_secret" "secret" {
  name = local.secret_name
}

module "diff_secret" {
  source                     = "./module"
  aws_secret_id_current      = data.aws_secretsmanager_secret.secret.id
  secret_manager_new_version = local.application_environment
  aws_region                 = local.aws_region
  aws_assume_role            = local.aws_assume_role
  aws_allowed_account_ids    = local.aws_allowed_account_ids
}
