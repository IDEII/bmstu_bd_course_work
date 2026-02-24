import { Component, Input } from '@angular/core';
import { FormGroup, FormBuilder } from '@angular/forms';
import { Tournament } from '../tournament';
import { TournamentService } from '../tournament.service';
import { RequestManagerService } from '../request-manager.service';
import { RequestWithTournamentTitle } from '../request';

@Component({
  selector: 'app-request-on-tournament',
  templateUrl: './request-on-tournament.component.html',
  styleUrls: ['./request-on-tournament.component.css']
})
export class RequestOnTournamentComponent {
  @Input() fighterid: string = '';
  @Input() category: number = 0;
  tournaments: Tournament[] = [];
  requestForm: FormGroup;
  fighterRequests: RequestWithTournamentTitle[] = []; 
  selectedRequestId: string | null = null;

  constructor(private fb: FormBuilder, private tournamentService: TournamentService, private requestManagerService: RequestManagerService) {
    this.tournaments = new Array<Tournament>();

    this.requestForm = this.fb.group({
      selectedTournament: ['']
    });
  }

  ngOnInit(): void {
    this.tournamentService.getTournaments().subscribe({
      next: (data: Tournament[]) => this.tournaments = data
    });
    this.loadRequests();

  }

  loadRequests() {
    this.requestManagerService.getRequestsByFighterId(this.fighterid).subscribe({
      next: (data: RequestWithTournamentTitle[]) => this.fighterRequests = data
    });
  }

  onSubmit(): void {
    const selectedTournamentId = this.requestForm.value.selectedTournament;
    const selectedTournament = this.tournaments.find(t => t.id === selectedTournamentId);

    if (this.fighterid && selectedTournament) {
        if (selectedTournament.category === this.category) {
            this.tournamentService.sendRequest(this.fighterid, selectedTournamentId).subscribe({
                next: (response: any) => {
                    console.log('Request sent successfully', response);
                    this.loadRequests();
                },
                error: (err: any) => console.error('Error sending request', err)
            });
        } else {
            console.error('Категория турнира не совпадает с категорией бойца');
            alert('Ошибка: категория турнира не совпадает с категорией бойца');
        }
    } else {
        console.error('Fighter ID is missing or Tournament ID is invalid');
        alert('Ошибка: ID бойца отсутствует или ID турнира неверен');
    }
    this.loadRequests();

}

  deleteRequest() {
    if (this.selectedRequestId) {
      this.requestManagerService.deleteRequest(this.selectedRequestId).subscribe({
        next: (response: any) => {
          console.log('Request deleted successfully', response);
          this.loadRequests(); 
          this.selectedRequestId = null;
        },
        error: (err: any) => console.error('Error deleting request', err)
      });
    } else {
      console.error('No request selected for deletion');
    }
    this.loadRequests();
  }
}