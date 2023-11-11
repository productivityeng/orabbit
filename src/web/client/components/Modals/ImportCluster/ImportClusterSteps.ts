export enum FormSteps {
    INIT,
    POPULATE_ACCESS,
    SELECT_QUEUE_TO_IMPORT,
    SELECT_EXCHANGE_TO_IMPORT,
    SELECT_USER_TO_IMPORT,
    FINISH,
  }
  
export const StepOrder: {
    [key in FormSteps]: { Next: FormSteps; Previous: FormSteps };
  } = {
    [FormSteps.INIT]: {
      Next: FormSteps.POPULATE_ACCESS,
      Previous: FormSteps.INIT,
    },
    
    [FormSteps.POPULATE_ACCESS]: {
        Next: FormSteps.SELECT_QUEUE_TO_IMPORT,
        Previous: FormSteps.INIT,
      },
      
    [FormSteps.SELECT_QUEUE_TO_IMPORT]: {
        Next: FormSteps.SELECT_EXCHANGE_TO_IMPORT,
        Previous: FormSteps.POPULATE_ACCESS,
      },
      
    [FormSteps.SELECT_EXCHANGE_TO_IMPORT]: {
        Next: FormSteps.SELECT_USER_TO_IMPORT,
        Previous: FormSteps.SELECT_QUEUE_TO_IMPORT,
      },
      
    [FormSteps.SELECT_USER_TO_IMPORT]: {
        Next: FormSteps.FINISH,
        Previous: FormSteps.SELECT_EXCHANGE_TO_IMPORT,
      },
      
    [FormSteps.FINISH]: {
        Next: FormSteps.FINISH,
        Previous: FormSteps.SELECT_USER_TO_IMPORT,
      },

  };
