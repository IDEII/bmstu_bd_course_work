import { Component } from '@angular/core';
import { OrganazerService } from '../organazer.service'
import { Organazer } from '../organazer';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-organazer-list',
  templateUrl: './organazer-list.component.html',
  styleUrl: './organazer-list.component.css'
})
export class OrganazerListComponent {
  organazers: Organazer[] = [];
   
    constructor(private OrganazerService: OrganazerService) {
      this.organazers = new Array<Organazer>();
  }
    ngOnInit() {
      
      this.OrganazerService.getOrganazers().subscribe(data => {
        this.organazers= data.filter(organazer => organazer.id !== '00000000-0000-0000-0000-000000000000');
      });
    }
}
